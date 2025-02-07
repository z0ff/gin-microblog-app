<script lang="ts">
    import "../app.css"

    import { onMount } from 'svelte';
    import { logout, getMe } from "$lib/auth";
    import type { User } from "$lib/types";

    let isLoggedin = false;
    let userInfo: Promise<User | null> | undefined = undefined;

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

        console.log('isLoggedin: ', isLoggedin);

        // ログインしている場合はユーザー情報を取得する
        if (isLoggedin) {
            userInfo = getMe();
        }
    });

    const onClickLogoutButton = async () => {
        await logout();
        location.href = "/login";
    };
</script>

<div class="navbar">
    <div class="flex-1">
        <a href="/" class="btn btn-ghost text-xl">microblog</a>
    </div>
    <div class="flex-none">
        {#await userInfo}
            <p>loading...</p>
        {:then data}
            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="btn btn-ghost">{data?.display_name}</div>
                <ul tabindex="0" role="menu" class="shadow menu dropdown-content bg-base-100 rounded-box w-52">
                    <li>
                        <a>profile</a>
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