<script lang="ts">
    import "../app.css"

    import { onMount } from 'svelte';
    import { logout, getMe } from "$lib/auth";
    import type { User } from "$lib/types";
    import { isLoggedInStore, userStore, searchQueryStore } from "$lib/stores";

    let userInfo: User | null = null;

    onMount(async () => {
        // ログイン中のユーザー情報を取得
        userInfo = await getMe();

        // ログインしていない場合はログインページにリダイレクト
        if (!(location.pathname === "/login" || location.pathname === "/signup")) {
            if (userInfo === null) {
                location.href = "/login";
            }
        } else {
            // ログインしている場合はトップページにリダイレクト
            if (userInfo !== null) {
                location.href = "/";
            }
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

    // ユーザー検索ボタンをクリックしたときの処理
    const onClickUserSearchButton = () => {
        const query = (document.getElementById('search-query') as HTMLInputElement).value;
        // searchQueryStore.set(query);
        location.href = `/usersearch?query=${query}`;
    };
</script>

<div class="navbar">
    <div class="navbar-start">
        <a href="/" class="btn btn-ghost text-xl">microblog</a>
    </div>
    <div class="navbar-center">
        <div class="join">
            <div>
                <div>
                    <input id="search-query" class="input input-bordered join-item" placeholder="検索したい文字列を入力" bind:value="{$searchQueryStore}" />
                </div>
            </div>
            <div class="indicator">
                <button on:click={onClickSearchButton} class="btn join-item">検索</button>
            </div>
            <div class="indicator">
                <button on:click={onClickUserSearchButton} class="btn join-item">ユーザー検索</button>
            </div>
        </div>
    </div>
    <div class="navbar-end">
        {#await userInfo}
            <p>loading...</p>
        {:then data}
            {#if data === null}
                <a href="/login" class="btn btn-ghost">ログイン</a>
            {:else}
            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="btn btn-ghost">{data?.display_name}</div>
                <ul tabindex="0" role="menu" class="shadow menu dropdown-content bg-base-100 rounded-box w-52">
                    <li>
                        <a href="/@{userInfo?.name}">プロフィール</a>
                    </li>
<!--                    クライアント側設定画面は一旦実装しない-->
<!--                    <li>-->
<!--                        <a>settings</a>-->
<!--                    </li>-->
                    <li>
                        <a on:click={onClickLogoutButton}>ログアウト</a>
                    </li>
                </ul>
            </div>
            {/if}
        {:catch error}
            <p>{error}</p>
        {/await}
    </div>
</div>

<div class="container max-w-3xl mx-auto">
    <slot />
</div>

<style>
    :global(body) {
        overflow-wrap: anywhere; /* 収まらない場合に折り返す */
        word-break: normal; /* 単語の分割はデフォルトに依存 */
        line-break: strict; /* 禁則処理を厳格に適用 */
    }
</style>