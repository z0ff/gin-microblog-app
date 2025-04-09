import { type Writable, writable } from "svelte/store";
import type { User } from "./types";

export const isLoggedInStore: Writable<boolean> = writable(false);
export const userStore: Writable<User | null> = writable(null);
export const searchQueryStore: Writable<string> = writable('');
