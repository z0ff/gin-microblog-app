<script lang="ts">
	import { onMount } from 'svelte';
	import { strToDate, utcToLocalTime } from "$lib/datetime";


	let postContent: HTMLTextAreaElement;
	let userId = 1;

	let postButtonEnabled = false;

	let postCharsCount = 0;

	onMount(() => {
		postContent = document.getElementById('post-content') as HTMLTextAreaElement;
	});

	const getPosts = async () => {
		const responseJson = await getPostsJson();
		if (responseJson.error) {
			return { error: responseJson.error };
		}

		const posts = JSON.parse(responseJson);
		// console.log(posts);

		return posts;
	};

	const getPostsJson = async () => {
		const res = await fetch('http://localhost:3000/', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
			},
		}).catch(console.error);

		if (res) {
			return await res.json()
		} else {
			return { error: res.statusText };
		}
	};
	let posts = getPostsJson();
	console.log(posts);

	const submitPost = async () => {
		const res = await fetch('http://localhost:3000/post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ user_id: userId, content: postContent.value.trim() }),
		}).catch(console.error);

		if (res) {
			console.log(await res.json());
		} else {
			console.error(res);
		}

		postContent.value = "";
		countPostContentChars();

		posts = await getPostsJson();
	};

	const toggleEnabledPostButton = () => {
		postButtonEnabled = postCharsCount > 0 && postCharsCount <= 500 && postContent.value.trim().length;
	};

	const countPostContentChars = () => {
		postCharsCount = postContent.value.length;
	}

	const onInput = () => {
		countPostContentChars();
		toggleEnabledPostButton();
	}


</script>

<textarea on:input={onInput} rows="10" cols="50" placeholder="type something here" id="post-content"></textarea>
<span>{postCharsCount}/500</span>
<button on:click={submitPost} type="button" id="post-submit" disabled={!postButtonEnabled}>submit</button>

{#await posts}
	<p>loading...</p>
{:then data}
	<div class="flex flex-col gap-y-1 max-w-3xl">
	{#each data as post}
		<div class="border-solid border-2 border-sky-500">
			<p class="font-bold">{post.User.DisplayName}</p>
			<p>{strToDate(post.CreatedAt)}</p>
			<p>{post.Content}</p>
		</div>
	{/each}
	</div>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
