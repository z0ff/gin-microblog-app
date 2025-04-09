<script lang="ts">
    import {getMe} from "$lib/auth";
    import { getUser, follow, unfollow, isFollowing } from "$lib/users";
    import { getPostsOfUser } from "$lib/posts";
    import type { PageData } from "./$types";
    import type {User} from "$lib/types";
    import {onMount} from "svelte";
    import Post from "$lib/Components/Post.svelte";

    export let data: PageData;

    const username = data.props.username;

    let user: Promise<User | null>;
    let posts: Promise<any>;

    const getPosts = async () => {
        let posts: any;
        const userdata = await user;
        if (userdata === null) {
            return { error: "user not found" };
        }
        posts = await getPostsOfUser(userdata.name);
        return posts;
    };

    user = getUser(username);
    posts = getPosts();

    onMount(async () => {
        const me = await getMe();
        const userdata = await user;
        const meActionsElem = document.getElementById('me-actions');
        const notmeActionsElem = document.getElementById('notme-actions');
        const followingActionsElem = document.getElementById('following-actions');

        if (me === null || userdata === null) {
            return;
        }

        if (meActionsElem === null || notmeActionsElem === null || followingActionsElem === null) {
            return;
        }

        if (me.name === userdata.name) {
            meActionsElem.classList.remove('hidden');
        } else {
            if (await isFollowing(userdata.name)) {
                followingActionsElem.classList.remove('hidden');
            } else {
                notmeActionsElem.classList.remove('hidden');
            }
        }
    });

    const onClickFollowButton = async () => {
        const userdata = await user;
        if (userdata === null) {
            return;
        }

        if (await follow(userdata.name)) {
            const followingActionsElem = document.getElementById('following-actions');
            const notmeActionsElem = document.getElementById('notme-actions');
            if (followingActionsElem === null || notmeActionsElem === null) {
                return;
            }
            followingActionsElem.classList.remove('hidden');
            notmeActionsElem.classList.add('hidden');
        } else {
            console.log('failed to follow');
        }
    };

    const onClickUnfollowButton = async () => {
        const userdata = await user;
        if (userdata === null) {
            return;
        }

        if (await unfollow(userdata.name)) {
            const followingActionsElem = document.getElementById('following-actions');
            const notmeActionsElem = document.getElementById('notme-actions');
            if (followingActionsElem === null || notmeActionsElem === null) {
                return;
            }
            followingActionsElem.classList.add('hidden');
            notmeActionsElem.classList.remove('hidden');
        } else {
            console.log('failed to unfollow');
        }
    };
</script>

{#await user}
    <p>loading...</p>
{:then data}
    <div class="card card-bordered shadow-lg max-w-3xl">
        <div class="card-body">
            <p class="text-2xl">{data?.display_name}</p>
            <p class="text-lg">@{data?.name}</p>
            <div id="me-actions" class="hidden">
                <button class="btn btn-primary">プロフィールを編集</button>
            </div>
            <div id="notme-actions" class="hidden">
                <button on:click={onClickFollowButton} class="btn btn-primary btn-outline">フォロー</button>
            </div>
            <div id="following-actions" class="hidden">
                <button on:click={onClickUnfollowButton} class="btn btn-primary">フォロー解除</button>
            </div>
        </div>
    </div>
{:catch error}
    <p style="color: red">{error.message}</p>
{/await}

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
