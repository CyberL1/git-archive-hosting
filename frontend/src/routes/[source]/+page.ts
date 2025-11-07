import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params: { source } }) => {
  const res = await fetch(`/api/repos/${source}`);
  const data = await res.json();
  return { repositories: data };
};
