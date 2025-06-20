#!/bin/bash
set -euo pipefail

for cmd in awk date aws; do
	if ! command -v "$cmd" > /dev/null 2>&1; then
		echo "Required command '$cmd' not found in PATH. Exiting..."
		exit 1
	fi
done

# Make sure all arguments are passed.
if [ "$#" -lt 2 ]; then
	echo "Usage: ${0} <bucket_name> <directory> [days_old]"
	echo "Example: ${0} example-bucket db_backup/ 7"
	exit 1
fi

BUCKET_NAME="$1"
DIRECTORY="$2"
DAYS_OLD="${3:-30}" # Default 30 days.
DRY_RUN="${4:-"true"}" # Dry run 'true' by default.

if [ "$DRY_RUN" = "true" ]; then
	echo "+--------- DRY RUN ----------+"
	echo "| No files will be deleted.  |"
	echo "+----------------------------+"
fi

# Get the desired max age date in UNIX seconds.
THRESHOLD_DATE=$(date -d "${DAYS_OLD} days ago" +%s)

OBJECT_LIST=$(aws ls "s3://${BUCKET_NAME}/${DIRECTORY}" || true)

# If no objects in bucket directory, exit pre-maturely.
if [ -z "${OBJECT_LIST}" ]; then
	echo "No objects found in s3://${BUCKET_NAME}/${DIRECTORY}. Exiting..."
	exit 0
fi

# List all objects in the desired bucket directory, and read through the output
# entires. 
echo ${OBJECT_LIST} | while read -r line; do
FILE_PATH=$(echo "$line" | awk '{print $4}')
echo "[FOUND] ${FILE_PATH}"

# Skip empty file names, i.e. directories.
if [ -z "$FILE_PATH" ]; then
	continue
fi

# Get the 'last modified' date in UNIX seconds for the object entry.
LAST_MODIFIED_STR=$(echo "$line" | awk '{print $1"T"$2}')
LAST_MODIFIED_SECONDS=$(date -d "${LAST_MODIFIED_STR}" +%s)

if [ "$LAST_MODIFIED_SECONDS" -lt "$THRESHOLD_DATE" ]; then
	URL="s3://${BUCKET_NAME}/${FILE_PATH}"
	if [ "$DRY_RUN" = "true" ]; then
		echo "(Dry Run) [DELETE] ${URL}"
	else
		echo "[DELETE] ${URL}"
		aws s3 rm ${URL}
	fi
fi
done;
