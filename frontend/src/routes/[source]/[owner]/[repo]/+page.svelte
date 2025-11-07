<script>
  import { page } from "$app/state";
</script>

Repo full name: {page.data.repository.source}/{page.data.repository.owner}/{page
  .data.repository.name}<br />
Created at: {page.data.repository.createdAt}<br /><br />
Clone command: git clone http://localhost:8080/{page.data.repository
  .source}/{page.data.repository.owner}/{page.data.repository.name}.git
