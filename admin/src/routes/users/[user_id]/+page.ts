import type { PageLoad } from './$types';
import {getUserDetails} from "$lib/users";
import type {User} from "$lib/types";

export const load: PageLoad = async ({ params }) => {
    let userId: number = parseInt(params.user_id);

    return {
        props: {
            user_id: userId,
        }
    }
}