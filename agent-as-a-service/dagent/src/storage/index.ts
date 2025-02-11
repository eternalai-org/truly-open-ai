import { LocalStorage } from "node-localstorage";

const storage_path = './db';
export const twitterStorage = new LocalStorage(storage_path + '/twitter-service-db');

export const postedStorage = new LocalStorage(storage_path + '/posted-db');
