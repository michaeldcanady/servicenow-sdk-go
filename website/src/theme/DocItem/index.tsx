import React from 'react';
import DocItem from '@theme-original/DocItem';
import type {ComponentProps} from 'react';
import DocsFeedback from '@site/src/components/DocsFeedback';

// Appends the per-page feedback widget to every docs page.
export default function DocItemWrapper(
  props: ComponentProps<typeof DocItem>,
): React.ReactElement {
  return (
    <>
      <DocItem {...props} />
      <DocsFeedback />
    </>
  );
}
