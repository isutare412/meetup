<script lang="ts">
  import DarkModeToggle from "./DarkModeToggle.svelte";
</script>

<div class="bg-base-100 shadow-md shadow-neutral relative">
  <div class="navbar bg-base-100 max-w-5xl mx-auto">
    <div class="navbar-start">
      <div class="dropdown">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label for="my-drawer" class="btn btn-ghost lg:hidden">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            ><path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h8m-8 6h16"
            /></svg
          >
        </label>
      </div>
      <a class="btn btn-ghost normal-case text-xl">Redshore</a>
    </div>
    <div class="navbar-center hidden lg:flex">
      <ul class="menu menu-horizontal p-0">
        <li><a>Item 1</a></li>
        <li tabindex="0">
          <a>
            Parent
            <svg
              class="fill-current"
              xmlns="http://www.w3.org/2000/svg"
              width="20"
              height="20"
              viewBox="0 0 24 24"
              ><path
                d="M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z"
              /></svg
            >
          </a>
          <ul class="p-2">
            <li><a>Submenu 1</a></li>
            <li><a>Submenu 2</a></li>
          </ul>
        </li>
        <li><a>Item 3</a></li>
      </ul>
    </div>
    <div class="navbar-end">
      <DarkModeToggle />
    </div>
  </div>
</div>
