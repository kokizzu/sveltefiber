import { writable } from 'svelte/store';

export let endpointUrl = (window && window.location.host === '127.0.0.1:5500') ? 'http://localhost:3001' : '';

export const listNames = writable([]);
