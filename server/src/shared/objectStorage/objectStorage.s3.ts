import { DeleteObjectCommand, GetObjectCommand, PutObjectCommand, S3Client } from "@aws-sdk/client-s3";
import type { ObjectStorage } from "./objectStorage";
import type { Readable } from "stream";

const REGIONS = [
  "us-east-2",
  "us-east-1",
  "us-west-1",
  "us-west-2",
  "af-south-1",
  "ap-east-1",
  "ap-south-2",
  "ap-southeast-3",
  "ap-southeast-5",
  "ap-southeast-4",
  "ap-south-1",
  "ap-northeast-3",
  "ap-northeast-2",
  "ap-southeast-1",
  "ap-southeast-2",
  "ap-northeast-1",
  "ca-central-1",
  "ca-west-1",
  "eu-central-1",
  "eu-west-1",
  "eu-west-2",
  "eu-south-1",
  "eu-west-3",
  "eu-south-2",
  "eu-north-1",
  "eu-central-2",
  "il-central-1",
  "me-south-1",
  "me-central-1",
  "sa-east-1",
  "us-gov-east-1",
  "us-gov-west-1",
] as const;


type S3Configuration = {
  bucket: string;
  region: typeof REGIONS[number]
}

export class S3ObjectStorage implements ObjectStorage {
  private s3: S3Client;
  private bucket: string
  private region: typeof REGIONS[number]

  constructor(cfg: S3Configuration) {
    const { bucket, region } = cfg
    this.bucket = bucket
    this.region = region

    this.s3 = new S3Client({ region })
  }

  uploadObject = async (key: string, data: Buffer): Promise<string> => {
    await this.s3.send(new PutObjectCommand({
      Bucket: this.bucket,
      Key: key,
      Body: data,
    }))

    return this.getObjectUrl(key)
  }

  getObject = async (key: string): Promise<ReadableStream> => {
    const res = await this.s3.send(new GetObjectCommand({
      Bucket: this.bucket,
      Key: key,
    }))

    if (!res.Body) {
      throw new Error(`Object with key ${key} not found`)
    }

    return res.Body.transformToWebStream()
  }

  deleteObject = async (key: string): Promise<void> => {
    await this.s3.send(new DeleteObjectCommand({
      Bucket: this.bucket,
      Key: key,
    }))
  }

  getObjectUrl = (key: string): string => {
    return `https://${this.bucket}.s3${this.region}.amazonaws.com/${key}`
  }
}
