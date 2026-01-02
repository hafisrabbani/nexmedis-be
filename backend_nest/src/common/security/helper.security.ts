import { randomBytes, createHash } from 'crypto';

/**
 * Generate random API key (64 hex chars)
 * Equivalent to Go: rand.Read + hex.EncodeToString
 */
export function generateApiKey(): string {
  return randomBytes(32).toString('hex');
}

/**
 * Hash API key using SHA-256
 * Equivalent to Go: sha256.Sum256 + hex.EncodeToString
 */
export function hashApiKey(key: string): string {
  return createHash('sha256').update(key).digest('hex');
}

/**
 * Encrypt email (simple deterministic encoding)
 * Equivalent to Go: []byte(email + ":" + APP_SECRET)
 *
 * NOTE:
 * This is NOT encryption, only deterministic obfuscation.
 * Matches original Go semantics exactly.
 */
export function encryptEmail(email: string): Buffer {
  const secret = process.env.APP_SECRET!;
  return Buffer.from(`${email}:${secret}`, 'utf-8');
}
