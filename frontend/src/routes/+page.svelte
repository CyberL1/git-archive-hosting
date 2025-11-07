<script>
  import { page } from "$app/state";
</script>

<h1>Repositories</h1>
{#each page.data.repositories as repository}
  <div>
    <h2>
      <a href="/{repository.source}">{repository.source}</a>/<a
        href="/{repository.source}/{repository.owner}">{repository.owner}</a
      >/<a href="/{repository.source}/{repository.owner}/{repository.name}">
        {repository.name}
      </a>
    </h2>
    <a href="/{repository.source}/{repository.owner}/{repository.name}">
      View Repository
    </a>
    <a href={repository.originalUrl} target="_blank">View Original</a>
  </div>
{/each}
