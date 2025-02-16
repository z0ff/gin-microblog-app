<script lang="ts">
    import {getMe, getUser} from "$lib/auth";
    import type { PageData } from "./$types";
    import type {User} from "$lib/types";
    import {onMount} from "svelte";
    import Post from "$lib/Components/Post.svelte";

    export let data: PageData;

    const username = data.props.username;

    let user: Promise<User | null>;
    let posts: Promise<any>;


    const getPostsJson = async (username: string) => {
        if (username===null) {
            return { error: "user not found" };
        }
        const res = await fetch(`http://localhost:3000/post/${username}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch((error) => {
            console.error(error);
            return;
        });

        if (res) {
            return await res.json()
        } else {
            return { error: res.statusText };
        }

    };
    const getPosts = async () => {
        let posts: any;
        const userdata = await user;
        if (userdata === null) {
            return { error: "user not found" };
        }
        posts = await getPostsJson(userdata.name);
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
            const isFollowing = await fetch(`http://localhost:3000/is_following/${userdata.name}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                // fetch APIはデフォルトでセッション情報を送信しないため、
                // 明示的に送信するように設定する
                credentials: "include"
            }).then(res => res.json()).catch(console.error);

            console.log(isFollowing.is_following);

            if (isFollowing.is_following) {
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
        const res = await fetch(`http://localhost:3000/follow/${userdata.name}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch(console.error);

        if (res) {
            console.log(await res.json());
        } else {
            console.error(res);
        }

        const followingActionsElem = document.getElementById('following-actions');
        const notmeActionsElem = document.getElementById('notme-actions');
        if (followingActionsElem === null || notmeActionsElem === null) {
            return;
        }
        followingActionsElem.classList.remove('hidden');
        notmeActionsElem.classList.add('hidden');
    };

    const onClickUnFollowButton = async () => {
        const userdata = await user;
        if (userdata === null) {
            return;
        }
        const res = await fetch(`http://localhost:3000/unfollow/${userdata.name}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            // fetch APIはデフォルトでセッション情報を送信しないため、
            // 明示的に送信するように設定する
            credentials: "include"
        }).catch(console.error);

        if (res) {
            console.log(await res.json());
        } else {
            console.error(res);
        }

        const followingActionsElem = document.getElementById('following-actions');
        const notmeActionsElem = document.getElementById('notme-actions');
        if (followingActionsElem === null || notmeActionsElem === null) {
            return;
        }
        followingActionsElem.classList.add('hidden');
        notmeActionsElem.classList.remove('hidden');
    };
</script>

{#await user}
    <p>loading...</p>
{:then data}
    <p>{data?.display_name}</p>
    <div id="me-actions" class="hidden">
        <button class="btn btn-primary">Edit Profile</button>
    </div>
    <div id="notme-actions" class="hidden">
        <button on:click={onClickFollowButton} class="btn btn-primary">Follow</button>
    </div>
    <div id="following-actions" class="hidden">
        <button on:click={onClickUnFollowButton} class="btn btn-primary">UnFollow</button>
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
