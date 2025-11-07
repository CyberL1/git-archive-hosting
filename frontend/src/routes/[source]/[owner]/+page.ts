import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params: { source, owner } }) => {
  const res = await fetch(`/api/repos/${source}/${owner}`);
  const data = await res.json();
  return { repositories: data };
};
