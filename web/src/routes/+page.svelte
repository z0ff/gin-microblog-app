<script lang="ts">
	import { onMount } from 'svelte';
	import { getTimeline, submitPost } from '$lib/posts';
	import Post from "$lib/Components/Post.svelte";

	let postContent: HTMLTextAreaElement;

	let postButtonEnabled = false;

	let postCharsCount = 0;

	onMount(() => {
		postContent = document.getElementById('post-content') as HTMLTextAreaElement;
	});

	let posts = getTimeline();

	const onClickSubmitBtn = async () => {
		if (await submitPost(postContent.value.trim())) {
			postContent.value = "";
			countPostContentChars();

			posts = await getTimeline();
		} else {
			alert("Failed to submit post");
		}
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

<div class="max-w-3xl my-4 mx-auto">
	<textarea rows="5" on:input={onInput} placeholder="今なにしてる？" id="post-content" class="textarea textarea-bordered textarea-primary w-full"></textarea>
	<div class="flex justify-between items-center">
		<span class="text-sm text-gray-500">{postCharsCount}/500</span>
		<button on:click={onClickSubmitBtn} type="button" id="post-submit" disabled={!postButtonEnabled} class="btn btn-primary">投稿</button>
	</div>
</div>

{#await posts}
	<p>loading...</p>
{:then data}
	<div class="flex flex-col gap-y-1 max-w-3xl">
	{#each data as post}
		<Post post={post} />
	{/each}
	</div>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
