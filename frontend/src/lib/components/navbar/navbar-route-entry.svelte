<script lang="ts">
	import { page } from '$app/state';

	type Props = {
		pathname: string;
		name: string;
	};

	let { pathname, name }: Props = $props();
</script>

<li>
	<a
		href={pathname}
		title={name}
		class:is-current={page.url.pathname === pathname}
		class="hover:text-text text-text/75 [.is-current]:text-text transition-colors before:invisible before:block before:h-0 before:overflow-hidden before:font-medium before:content-[attr(title)] [.is-current]:font-medium"
	>
		{name}
	</a>
</li>
