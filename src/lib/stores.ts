import { type Writable, writable } from "svelte/store";

export const searchQueryStore: Writable<string> = writable('');
