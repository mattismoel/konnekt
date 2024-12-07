import { env } from "@/config/env";
import type { ResultSet } from "@libsql/client";
import type { ExtractTablesWithRelations } from "drizzle-orm";
import { drizzle } from "drizzle-orm/libsql";
import type { SQLiteTransaction } from "drizzle-orm/sqlite-core";

export type TX = SQLiteTransaction<"async", ResultSet, Record<string, never>, ExtractTablesWithRelations<Record<string, never>>>
export const db = drizzle(env.DSN)
