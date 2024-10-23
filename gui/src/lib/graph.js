/**
  * @typedef ProjectFile
  * @type {object}
  * @property {string} packageName
  * @property {string} qualifiedPackageName
  * @property {string} filePath
  * @property {string} modPath
  * @property {string[]} imports
  * @property {string[]} externalImports
  */

/**
  * @param {import('graphology').default} graph
  * @param {Object.<string, ProjectFile>} json
*/
export const ReadGraph = (graph, json) => {
  let sz = 0;
  for (const [filePath, file] of Object.entries(json)) {
    graph.addNode(filePath, {
      label: file.qualifiedPackageName,
      x: sz,
      y: sz++,
      size: 15,
      color: 'red',
    });
  }
  for (const [filePath, file] of Object.entries(json)) {
    console.log(file);

    for (const imp of file.imports) {
      if (json[imp]) {
        graph.addEdge(filePath, imp, {
          size: 5,
          color: 'black',
        });
      } else {
        console.warn('Missing import edge:', imp);
      }
    }
  }
};
