import React, {useEffect, useState} from 'react';
import {
  DEFAULT_DISPLAY,
  PlaceholderValues,
  readStoredValues,
  storeValues,
} from '../GoSnippet/placeholders';
import styles from './styles.module.css';

// Lets the reader enter their own instance details once; every snippet-driven
// code sample on the site then renders with those values. Values live only in
// the browser's localStorage.

const FIELDS: Array<{key: keyof PlaceholderValues; label: string; hint: string}> = [
  {key: 'instance', label: 'Instance', hint: 'dev12345'},
  {key: 'username', label: 'Username', hint: 'integration.user'},
  {key: 'table', label: 'Table', hint: 'incident'},
  {key: 'sysId', label: 'Record sys_id', hint: ''},
];

export default function TryItValues(): React.ReactElement {
  const [values, setValues] = useState<PlaceholderValues>({});

  useEffect(() => {
    setValues(readStoredValues());
  }, []);

  const update = (key: keyof PlaceholderValues, value: string) => {
    const next = {...values, [key]: value};
    setValues(next);
    storeValues(next);
  };

  return (
    <details className={styles.card}>
      <summary className={styles.summary}>
        Personalize the code samples
        <span className={styles.badge}>stays in your browser</span>
      </summary>
      <p className={styles.note}>
        Enter your instance details and every code sample in the docs renders
        with them instead of placeholders like{' '}
        <code>{DEFAULT_DISPLAY.instance}</code>. Values are stored only in this
        browser's localStorage — never sent anywhere. Don't enter a password;
        samples keep the <code>{DEFAULT_DISPLAY.password}</code> placeholder.
      </p>
      <div className={styles.grid}>
        {FIELDS.map(({key, label, hint}) => (
          <label key={key} className={styles.field}>
            <span className={styles.label}>{label}</span>
            <input
              type="text"
              value={values[key] ?? ''}
              placeholder={hint}
              onChange={(e) => update(key, e.target.value)}
            />
          </label>
        ))}
      </div>
    </details>
  );
}
