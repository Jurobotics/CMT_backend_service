import { DB, QueryParameterSet, Row, RowObject } from "sqlite";
import * as path from "path";

class Database {
  private db: DB;
  protected dir = path.join(Deno.cwd().split("src")[0], "data");

  constructor() {
    this.db = new DB();
  }

  /**
   * @param file SQL File
   * @returns
   */
  public async init(file: string): Promise<string> {
    this.db = new DB(path.join(this.dir, "CMT.db"));
    const fileExtension = file.split(".").pop();
    if (fileExtension !== "sql") {
      throw new Error("File is not a SQL file");
    }

    try {
      const sql = await Deno.readTextFile(path.join(this.dir, file));
      await this.db.execute(sql);
      return `Database initialized with file: ${file}`;
    } catch (error) {
      throw new Error(error);
    }
  }

  public async close(): Promise<void> {
    await this.db.close();
  }

  public async query(sql: string, args: QueryParameterSet = []): Promise<Row> {
    const res = await this.db.query(sql, args);
    return res;
  }

  public async queryEntries(
    sql: string,
    args: QueryParameterSet = []
  ): Promise<Array<any>> {
    const res = await this.db.queryEntries<any>(sql, args);
    return res;
  }
}

// TODO
export abstract class DatabaseEntity {
  abstract __create(): boolean;
  abstract __edit(): boolean;
  abstract __delete(): boolean;
}

export const db = new Database();
