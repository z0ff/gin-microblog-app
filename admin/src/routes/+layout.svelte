<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div class="navbar bg-base-100 shadow-sm">
  <div class="flex-1">
    <a class="btn btn-ghost text-xl" href="/">管理画面</a>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal px-1">
      <li><a href="/users">ユーザー管理</a></li>
      <li><a href="/posts">投稿管理</a></li>
    </ul>
  </div>
  <div class="flex-none">
    <a class="btn btn-ghost" href="/logout">ログアウト</a>
  </div>
</div>

<div class="mx-auto max-w-7xl p-4">
  {@render children()}
</div>
