<script lang="ts">
    import { page } from '$app/stores';
    import { searchUser } from "$lib/users";
    import { searchQueryStore } from "$lib/stores";
    import type {User} from "$lib/types";

    let users: Promise<User[]>;
    // 検索クエリパラメータ
    const query = $page.url.searchParams.get('query')
    searchQueryStore.set(query);

    users = searchUser(query);
    console.log(users);
</script>

{#await users}
    <p>loading...</p>
{:then data}
    <div class="flex flex-col gap-y-1 max-w-3xl">
        {#each data as user}
            <div class="card card-bordered shadow-lg max-w-3xl">
                <a href="/@{user.name}">
                    <div class="card-body">
                        <p class="text-2xl">{user.display_name}</p>
                        <p class="text-lg">@{user.name}</p>
                    </div>
                </a>
            </div>
        {:else}
            <p>no users</p>
        {/each}
    </div>
{:catch error}
    <p style="color: red">{error.message}</p>
{/await}
