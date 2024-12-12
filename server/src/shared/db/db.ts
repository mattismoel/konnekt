import { env } from "@/config/env";
import { createClient, type ResultSet } from "@libsql/client";
import type { ExtractTablesWithRelations } from "drizzle-orm";
import { drizzle } from "drizzle-orm/libsql";
import type { SQLiteTransaction } from "drizzle-orm/sqlite-core";

export type TX = SQLiteTransaction<"async", ResultSet, Record<string, never>, ExtractTablesWithRelations<Record<string, never>>>
const client = createClient({ url: env.DSN })
export const db = drizzle({ client })
