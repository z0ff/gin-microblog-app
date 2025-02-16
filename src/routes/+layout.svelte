<script lang="ts">
    import "../app.css"

    import { onMount } from 'svelte';
    import { logout, getMe } from "$lib/auth";
    import type { User } from "$lib/types";
    import { isLoggedInStore, userStore, searchQueryStore } from "$lib/stores";

    let isLoggedin = false;
    // let userInfo: Promise<User | null> | undefined = undefined;
    let userInfo: User | null = null;

    onMount(async () => {
        // ログインしているかどうかを確認する
        await fetch('http://localhost:3000/is_logged_in', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: "include"
        }).then(res => {
            // ログイン/会員登録画面以外の場合はログインしていない場合はログインページにリダイレクト
            if (location.pathname !== "/login" && location.pathname !== "/signup" && res.status === 401) {
                location.href = "/login";
            }
            // ログイン/会員登録画面の場合はログインしている場合はトップページにリダイレクト
            if ((location.pathname === "/login" || location.pathname === "/signup") && res.status === 200) {
                location.href = "/";
            }

            // ログインしているかどうかを取得する
            isLoggedin = res.status === 200;

            console.log('status: ', res.status);
        }).catch(console.error);

        isLoggedInStore.set(isLoggedin);

        console.log('isLoggedin: ', isLoggedin);

        // ログインしている場合はユーザー情報を取得する
        if (isLoggedin) {
            userInfo = await getMe();
            userStore.set(userInfo);
        }
    });

    const onClickLogoutButton = async () => {
        await logout();
        location.href = "/login";
    };

    // 検索ボタンをクリックしたときの処理
    const onClickSearchButton = () => {
        const query = (document.getElementById('search-query') as HTMLInputElement).value;
        // searchQueryStore.set(query);
        location.href = `/search?query=${query}`;
    };
</script>

<div class="navbar">
    <div class="flex-1">
        <a href="/" class="btn btn-ghost text-xl">microblog</a>
    </div>
    <div class="flex-1">
        <div class="join">
            <div>
                <div>
                    <input id="search-query" class="input input-bordered join-item" placeholder="Search" bind:value="{$searchQueryStore}" />
                </div>
            </div>
            <div class="indicator">
                <button on:click={onClickSearchButton} class="btn join-item">Search</button>
            </div>
        </div>
    </div>
    <div class="flex-none">
        {#await userInfo}
            <p>loading...</p>
        {:then data}
            {#if data === null}
                <a href="/login" class="btn btn-ghost">login</a>
            {:else}
            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="btn btn-ghost">{data?.display_name}</div>
                <ul tabindex="0" role="menu" class="shadow menu dropdown-content bg-base-100 rounded-box w-52">
                    <li>
                        <a href="/@{userInfo?.name}">profile</a>
                    </li>
<!--                    クライアント側設定画面は一旦実装しない-->
<!--                    <li>-->
<!--                        <a>settings</a>-->
<!--                    </li>-->
                    <li>
                        <a on:click={onClickLogoutButton}>logout</a>
                    </li>
                </ul>
            </div>
            {/if}
        {:catch error}
            <p>{error}</p>
        {/await}
    </div>
</div>

<div class="container mx-auto">
    <slot />
</div>

<style>
    :global(body) {
        overflow-wrap: anywhere; /* 収まらない場合に折り返す */
        word-break: normal; /* 単語の分割はデフォルトに依存 */
        line-break: strict; /* 禁則処理を厳格に適用 */
    }
</style>