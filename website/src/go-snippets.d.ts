// Go snippet files are imported as raw source strings via the
// go-raw-source plugin in docusaurus.config.ts.
declare module '*.go' {
  const content: string;
  export default content;
}
