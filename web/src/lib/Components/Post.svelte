<script lang="ts">
    import { formatIsoDateStr, strToDate } from "$lib/datetime";
    import { like, unlike } from "$lib/posts";
    import type { Post } from "$lib/types";

    export let post: Post;

    const onLikeBtnClicked = async () => {
        if (await like(post.id)) {
            post.is_liked = true;
        } else {
            console.error("Failed to like");
        }
    }

    const onUnlikeBtnClicked = async () => {
        if (await unlike(post.id)) {
            post.is_liked = false;
        } else {
            console.error("Failed to unlike");
        }
    }
</script>
<div class="card card-bordered bg-base-50 shadow-lg">
    <div class="card-body">
        <div class="flex gap-2">
            <div class="flex-none">
                <a href="/@{post.user.name}"><span class="font-bold">{post.user.display_name}</span></a>
            </div>
            <div class="flex-1">
                <a href="/@{post.user.name}">@{post.user.name}</a>
            </div>
            <div class ="flex-none">
                <p>{formatIsoDateStr(post.createdAt, 'yyyy/MM/dd HH:mm')}</p>
            </div>
        </div>
        <p class="whitespace-pre-wrap">{post.content}</p>
        <button class="btn {post.is_liked ? 'btn-primary' : 'btn-ghost'}" on:click={post.is_liked ? onUnlikeBtnClicked : onLikeBtnClicked}>{post.is_liked ? "いいね解除" : "いいね"}</button>
    </div>
</div>
