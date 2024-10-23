<script>
	import Sigma from 'sigma';
	import Graph from 'graphology';
	import { onMount } from 'svelte';
	import { ReadGraph } from '$lib/graph';

	/**
	 * @type {HTMLElement}
	 */
	let container;

	let graph = new Graph();

	/**
	 * @type {FileList}
	 */
	let files;

	onMount(() => {
		const sigmaInstance = new Sigma(graph, container);
	});

	$: if (files) {
		const reader = new FileReader();
		reader.onload = () => {
            try {
                const json = JSON.parse(String(reader.result));
                ReadGraph(graph, json);
            } catch (e) {
                console.error(e);
            }
		};
		reader.readAsText(files[0]);
	}
</script>

<label for="file">Select a file:</label>
<input type="file" id="file" name="file" bind:files />
<div bind:this={container} style="width: 800px; height: 600px; background: white" />
