<script lang="ts">
    import { page } from '$app/stores';
    import { strToDate } from "$lib/datetime";
    import { searchQueryStore} from "$lib/stores";
    import Post from "$lib/Components/Post.svelte";

    let posts: Promise<any>;
    // 検索クエリパラメータ
    const query = $page.url.searchParams.get('query')
    searchQueryStore.set(query);

    const getPostsJson = async (query: string | null) => {
        const res = await fetch('http://localhost:3000/search?query=' + query, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch(console.error);

        if (res) {
            return await res.json()
        } else {
            return { error: res.statusText };
        }
    };
    posts = getPostsJson(query);
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
