import { redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params: { source, owner, repo } }) => {
  if (repo.endsWith(".git")) {
    redirect(307, `/${source}/${owner}/${repo.slice(0, -4)}`);
  }

  const res = await fetch(`/api/repos/${source}/${owner}/${repo}`);
  const data = await res.json();
  return { repository: data };
};
