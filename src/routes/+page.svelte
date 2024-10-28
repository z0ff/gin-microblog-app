<script lang="ts">
	import { onMount } from 'svelte';

	let response: Promise<any>;

	onMount(() => {
		response = (async () => {
			const res = await fetch('http://localhost:3000/', {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},
			}).catch(console.error);

			// return res.json();

			if (res) {
				return res.json();
			} else {
				return { error: res.statusText };
			}
		})();

		response.then(console.log);
	});


</script>

<h1>Welcome to SvelteKit</h1>
{#await response}
	<p>loading...</p>
{:then data}
	<pre>{JSON.stringify(data, null, 2)}</pre>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
