<script lang="ts">
    import { getPosts, deletePost } from "$lib/posts";

    const posts = getPosts();
</script>
<h1 class="text-2xl">投稿管理</h1>

<table class="table table-zebra">
    <thead>
    <tr>
        <th>ID</th>
        <th>ユーザーID</th>
        <th>本文</th>
        <th>作成日時</th>
        <th>操作</th>
    </tr>
    </thead>
    <tbody>
    {#await posts}
        <p>loading...</p>
    {:then data}
        {#each data as post}
            <tr>
                <td>{post.id}</td>
                <td><a href="/users/{post.user_id}">{post.user_id}</a></td>
                <td>{post.content}</td>
                <td>{post.createdAt}</td>
                <td>
                    {#if post.deletedAt}
                        <button class="btn btn-disabled">削除</button>
                    {:else}
                        <button class="btn btn-error" on:click="{() => {deletePost(post.id)}}">削除</button>
                    {/if}
                </td>
            </tr>
        {/each}
    {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
    </tbody>
</table>
