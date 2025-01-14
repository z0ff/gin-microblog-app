<script lang="ts">
	import { onMount } from 'svelte';


	let postContent: HTMLTextAreaElement;
	let userId = 1;

	let postButtonEnabled = false;

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
			body: JSON.stringify({ user_id: userId, content: postContent.value }),
		}).catch(console.error);

		if (res) {
			console.log(await res.json());
		} else {
			console.error(res);
		}

		posts = await getPostsJson();
	};

	const toggleEnabledPostButton = () => {
		postButtonEnabled = postContent.value.length > 0;
	};


</script>

<h1>microblog</h1>

<textarea on:input={toggleEnabledPostButton} rows="10" cols="50" placeholder="type something here" id="post-content"></textarea>
<button on:click={submitPost} type="button" id="post-submit" disabled={!postButtonEnabled}>submit</button>

{#await posts}
	<p>loading...</p>
{:then data}
	{#each data as post}
		<div>
			<p>by {post.User.DisplayName}</p>
			<p>{post.CreatedAt}</p>
			<p>{post.Content}</p>
		</div>
	{/each}
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
