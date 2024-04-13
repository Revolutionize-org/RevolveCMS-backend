import pg from 'pg';

import dotenv from 'dotenv';
import { fileURLToPath } from 'url';
import { dirname } from 'path';
import path from 'path';

import argon2 from 'argon2';
import { v4 as uuidv4 } from 'uuid';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
dotenv.config({ path: path.resolve(__dirname, '../../.env') });

const client = createClient();
await client.connect();

/**
 * Role ID
 * ('d7de28aa-5028-4bee-8361-7d630d86da54', 'admin'),
 * ('d44e3f29-0ab5-40d4-b5d0-1e41c3cc59d3', 'user');
 */
createUser(
  client,
  'admin',
  'admin@proton.me',
  'adminPassword',
  'd7de28aa-5028-4bee-8361-7d630d86da54',
  '45955517-30ee-4310-b253-d0cd677cc92e'
);

function createClient() {
  const { Client } = pg;
  return new Client({
    host: 'localhost',
    port: process.env.POSTGRES_PORT,
    user: process.env.POSTGRES_USER,
    password: process.env.POSTGRES_PASSWORD,
    database: process.env.POSTGRES_DB,
  });
}

async function createUser(client, name, email, password, role_id, website_id) {
  const uuid = uuidv4();
  const password_hash = await argon2.hash(password);

  const query = `INSERT INTO 
  users(id, name, email, password_hash, role_id, website_id) 
  VALUES($1, $2, $3, $4, $5, $6)`;

  try {
    await client.query(query, [
      uuid,
      name,
      email,
      password_hash,
      role_id,
      website_id,
    ]);
    console.log('User created successfully!');
  } catch (error) {
    console.error('Error creating user:', error);
  } finally {
    await client.end();
  }
}
