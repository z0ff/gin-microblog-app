<script lang="ts">
  import { getUsers, deleteUser } from "$lib/users";
    import type { User } from "$lib/types";

    let users: Promise<User[]>;

    users = getUsers();
</script>
<h1 class="text-2xl">ユーザー管理</h1>

<table class="table table-zebra">
  <thead>
    <tr>
      <th>ID</th>
      <th>ユーザー名</th>
      <th>ユーザー表示名</th>
      <th>作成日時</th>
      <th>更新日時</th>
        <th>ステータス</th>
        <th>操作</th>
    </tr>
  </thead>
  <tbody>
  {#await users}
    <p>loading...</p>
    {:then data}
    {#each data as user}
        <tr>
        <td>{user.id}</td>
        <td>{user.name}</td>
            <td>{user.display_name}</td>
            <td>{user.createdAt}</td>
            <td>{user.updatedAt}</td>
            <td>
                {user.deletedAt}
                {#if user.deletedAt}
                    削除済み
                {:else}
                    有効
                {/if}
            </td>
            <td>
<!--                <a class="btn btn-primary" href="/users/{user.id}">詳細</a>-->
                {#if user.deletedAt}
                    <button class="btn btn-disabled">削除</button>
                {:else}
                    <button class="btn btn-error" on:click="{() => {deleteUser(user.id)}}">削除</button>
                {/if}
            </td>
        </tr>
    {/each}
    {:catch error}
    <p style="color: red">{error.message}</p>
    {/await}
    </tbody>
</table>
