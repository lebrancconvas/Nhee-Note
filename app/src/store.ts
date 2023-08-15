import { writable } from 'svelte/store';

export const showModal = writable(false);

export const title = writable('');
export const description = writable('');
export const price = writable(0);
