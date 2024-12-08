import { sha256 } from "@oslojs/crypto/sha2";
import { encodeBase32NoPadding, encodeHexLowerCase } from "@oslojs/encoding";
import type { Session } from "./session.dto";
import { SESSION_LIFETIME_DAYS } from "./constant";
import { addDays } from "date-fns";
import type { Response } from "express";

/**
 * @description Generates a session token.
 */
export const generateSessionToken = (): string => {
  const bytes = new Uint8Array(20);
  crypto.getRandomValues(bytes);
  const token = encodeBase32NoPadding(bytes)
  return token
}

/**
 * @description Returns the session ID generated from the input token.
 */
export const sessionIDFromToken = (token: string) => {
  const sessionID = encodeHexLowerCase(sha256(new TextEncoder().encode(token)))
  return sessionID
}

/**
 * @description Returns a {@link Session} object.
 * @param {string} token - The token to base the session on.
 * @param {number} userID - The associated user's ID.
 */
export const createSession = async (token: string, userID: number): Promise<Session> => {
  const sessionID = sessionIDFromToken(token)

  const session: Session = {
    id: sessionID,
    userID,
    expiresAt: addDays(new Date(), SESSION_LIFETIME_DAYS)
  }

  return session
}

/**
 * @description Sets the token cookie on the request.
 * @param {Response} res - The response to put the session cookie in.
 * @param {string} name - The name of the session cookie.
 * @param {string} token - The session token of the cookie.
 * @param {Date} expiresAt - The expiry date of the session cookie.
 */
export const setSessionTokenCookie = (res: Response, name: string, token: string, expiresAt: Date) => {
  res.cookie(name, token, {
    httpOnly: true,
    sameSite: "lax",
    expires: expiresAt,
    path: "/"
  })
}

/**
 * @description Deletes the session cookie on the input request.
 * @param {Response} res - The response to delete the cookie from.
 * @param {string} name - The name of the session cookie.
 */
export const deleteSessionTokenCookie = (res: Response, name: string) => {
  res.cookie(name, "", {
    httpOnly: true,
    sameSite: "lax",
    maxAge: 0,
    path: "/"
  })
}
