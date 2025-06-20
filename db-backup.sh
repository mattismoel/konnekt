#!/bin/bash
set -euo pipefail

VOLUME_NAME=${1}
FILE_PATH=${2}
BUCKET_NAME=${3}
BUCKET_DIRECTORY=${4}

# On script exit, clean up container and delete directory.
cleanup() {
	[[ -n "$CONTAINER_ID" ]] && docker rm -f "${CONTAINER_ID}" > /dev/null 2>&1
	[[ -d "${TMP_DIR}" ]] && rm -rf "${TMP_DIR}"
}

trap cleanup EXIT


# Make sure that all needed commands are present.
for cmd in docker aws mktemp; do
	if ! command -v "$cmd" > /dev/null 2>&1; then
		echo "Required command '$cmd' not found in PATH. Exiting..."
		exit 1
	fi
done

# Make sure all arguments are passed.
if [ "$#" -lt 4 ]; then
	echo "Usage: ${0} <volume_name> <file_path> <bucket_name> <bucket_directory>"
	echo "Example: ${0} some_named_volume data.db example-bucket db-backup"
	exit 1
fi

# Make sure the provided docker volume exists.
if ! docker volume inspect "${VOLUME_NAME}" > /dev/null 2>&1; then
	echo "Volume '${VOLUME_NAME}' does not exist. Exiting..."
	exit 1
fi

TMP_DIR=$(mktemp -d)
CONTAINER_ID=$(docker run -d -v ${VOLUME_NAME}:/app/data busybox true)

# Copy the database file to host backup path.
echo "Copying ${FILE_PATH} from volume '${VOLUME_NAME}'..."
docker cp "${CONTAINER_ID}:/app/data/${FILE_PATH}" "${TMP_DIR}/backup.db"

if [! -f "${TMP_DIR}/backup.db" ]; then
	echo "Failed to cipy file '${FILE_PATH}'. Exiting..."
	exit 1
fi

# Generate date string. Example: 2025-12-31T14-30-21
DATE_STR=$(date +'%Y-%m-%dT%H-%M-%S')
FILENAME="db_backup_${DATE_STR}.db"
S3_URL="s3://${BUCKET_NAME}/${BUCKET_DIRECTORY}${FILENAME}"

echo "Uploading to ${S3_URL}..."
aws s3 cp "${TMP_DIR}/backup.db" ${S3_URL}

echo "Backup complete: ${S3_URL}"
