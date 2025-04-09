<script lang="ts">
    import { page } from '$app/stores';
    import { searchPosts } from "$lib/posts";
    import { searchQueryStore } from "$lib/stores";
    import Post from "$lib/Components/Post.svelte";

    let posts: Promise<any[]>;
    // 検索クエリパラメータ
    const query = $page.url.searchParams.get('query')
    searchQueryStore.set(query);

    posts = searchPosts(query);
    console.log(posts);
</script>

{#await posts}
    <p>loading...</p>
{:then data}
    <div class="flex flex-col gap-y-1 max-w-3xl">
        {#each data as post}
            <Post post={post} />
        {:else}
            <p>no posts</p>
        {/each}
    </div>
{:catch error}
    <p style="color: red">{error.message}</p>
{/await}
