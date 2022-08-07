<script lang="ts">
  import NavBar from "@component/NavBar.svelte";
  import SideBar from "@component/SideBar.svelte";
</script>

<SideBar>
  <svelte:fragment slot="navbar">
    <NavBar />
  </svelte:fragment>
  <svelte:fragment slot="content">
    <main>
      <div class="max-w-5xl mx-auto p-4 bg-base-200">
        <h1>Hello World!</h1>
        <button class="btn btn-primary">Button</button>
        <button class="btn btn-secondary">Button</button>
        <button class="btn btn-accent">Button</button>
      </div>
    </main>
  </svelte:fragment>
</SideBar>

<style>
</style>
