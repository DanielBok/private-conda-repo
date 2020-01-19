declare module "markdown-to-jsx" {
  export interface MarkdownOptions {
    // Forces all string to be "blocks" - header like objects
    forceBlock?: boolean;

    // The inverse of forceBlock. Forces all strings to be rendered as strings, no indents
    forceInline?: boolean;

    // Overrides any React JSX like objects in the markdown with the actual component
    overrides?: {
      [component: string]:
        | {
            component: React.ReactNode;
            props?: object;
          }
        | React.ReactNode;
    };

    // By default, a lightweight deburring function is used to generate an HTML id from headings.
    // You can override this by passing a function to options.slugify. This is helpful when you are
    // using non-alphanumeric characters (e.g. Chinese or Japanese characters) in headings.
    slugify?: (id: string) => string;

    // By default only a couple of named html codes are converted to unicode characters. These include
    // characters such as &, ', >, <, &nbsp and " which have special meaning in HTML or other context.
    // Some projects require to extend this map of named codes and unicode characters.
    // To customize this list with additional html codes pass the option namedCodesToUnicode
    // as object with the code names needed as in the example below:
    namedCodesToUnicode?: {
      [character: string]: string;
    };
  }

  export interface MarkdownProps {
    options?: MarkdownOptions;
  }

  export function compiler(
    markdown: string,
    options?: MarkdownOptions
  ): React.ReactNode;

  export default class Markdown extends React.PureComponent<MarkdownProps> {}
}
