<script setup>
import SnapcartLogo from '@/components/icons/Snapcart-Logo.vue';

</script>
<template>
    <div class="flex min-h-screen flex-col justify-center px-6 py-12 lg:px-8 bg-gray-900">
      <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <SnapcartLogo class="w-full flex justify-center scale-150"/>
        <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-200">Sign in to your account</h2>
      </div>
  
      <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form class="space-y-6" action="#">
          <div>
            <label for="email" class="block text-sm font-medium leading-6 text-gray-200">Email address</label>
            <div class="mt-2">
              <input id="email" name="email" type="email" autocomplete="email" required class="block w-full rounded-md border-0 p-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white">
            </div>
          </div>
      
          <div>
            <div class="flex items-center justify-between">
              <label for="password" class="block text-sm font-medium leading-6 text-gray-200">Password</label>
              <div class="text-sm">
                <a href="#" class="font-semibold text-sky-600 hover:text-sky-500">Forgot password?</a>
              </div>
            </div>
            <div class="mt-2">
              <input id="password" name="password" type="password" autocomplete="current-password" required class="block w-full rounded-md border-0 p-1.5 placeholder:text-gray-400 sm:text-sm sm:leading-6 bg-gray-800 ring-1 ring-gray-700 text-white">
            </div>
          </div>
      
          <div>
            <button type="submit" class="flex w-full justify-center rounded-md bg-sky-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-sky-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-600">Sign in</button>
          </div>
        </form>
    
        <p class="mt-10 text-center text-sm text-gray-500">
          Not a member?
          <RouterLink to="/register" class="font-semibold leading-6 text-sky-600 hover:text-sky-500">Register Now</RouterLink>
        </p>
      </div>
    </div>
</template>