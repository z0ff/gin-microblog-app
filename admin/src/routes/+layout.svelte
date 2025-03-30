<script lang="ts">
	import '../app.css';
    import { onMount } from 'svelte';
    import { logout, getMe } from "$lib/auth";
    import type { Admin } from "$lib/types";

    let { children } = $props();
    let me: Admin | null;

    onMount(async () => {
        document.title = '管理画面';

        // ログイン中のユーザー情報を取得l
        me = await getMe();

        console.log('me: ', me);

        // ログインしていない場合はログインページにリダイレクト
        if (!(location.pathname === "/login" || location.pathname === "/signup")) {
            if (me === null) {
              location.href = "/login";
            }
        } else {
            // ログインしている場合はトップページにリダイレクト
            if (me !== null) {
              location.href = "/";
            }
        }
    });

    // ログアウト処理
    const onClickLogoutButton = async () => {
        await logout();
        location.href = "/login";
    };
</script>

<div class="navbar bg-base-100 shadow-sm">
  <div class="flex-1">
    <a class="btn btn-ghost text-xl" href="/">管理画面</a>
  </div>
  {#await getMe() then user}
  {#if user !== null}
  <div class="flex-none">
    <ul class="menu menu-horizontal px-1">
      <li><a href="/users">ユーザー管理</a></li>
      <li><a href="/posts">投稿管理</a></li>
    </ul>
  </div>
  <div class="flex-none">
    <button class="btn btn-ghost" on:click={onClickLogoutButton}>ログアウト</button>
  </div>
    {:else}
    <div class="flex-none">
        <a class="btn btn-ghost" href="/login">ログイン</a>
    </div>
    {/if}
    {/await}
</div>

<div class="container mx-auto max-w-7xl p-4">
  {@render children()}
</div>
