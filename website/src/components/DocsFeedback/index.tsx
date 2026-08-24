import React, {useEffect, useState} from 'react';
import {useLocation} from '@docusaurus/router';
import {trackEvent} from '../../utils/analytics';
import styles from './styles.module.css';

type Vote = 'yes' | 'no';

// Per-page feedback widget. Votes go to Plausible as a "Docs Feedback"
// event with the page path and vote as props — Plausible itself is the
// backend, so no server-side component is needed.
export default function DocsFeedback(): React.ReactElement {
  const {pathname} = useLocation();
  const [vote, setVote] = useState<Vote | null>(null);

  // Reset when navigating between docs pages.
  useEffect(() => {
    setVote(null);
  }, [pathname]);

  const send = (value: Vote) => {
    if (vote === null) {
      trackEvent('Docs Feedback', {page: pathname, vote: value});
    }
    setVote(value);
  };

  return (
    <div className={styles.container}>
      {vote === null ? (
        <>
          <span className={styles.prompt}>Was this page helpful?</span>
          <button
            type="button"
            className="button button--sm button--outline button--secondary"
            onClick={() => send('yes')}>
            Yes
          </button>
          <button
            type="button"
            className="button button--sm button--outline button--secondary"
            onClick={() => send('no')}>
            No
          </button>
        </>
      ) : (
        <span className={styles.thanks}>Thanks for your feedback!</span>
      )}
    </div>
  );
}
