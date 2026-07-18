import {useEffect, useState} from 'react';

// User-editable placeholder values (see the TryItValues component). Stored in
// localStorage so every snippet on every page renders with the reader's own
// instance details. Falls back to neutral display values.

export type PlaceholderValues = {
  instance?: string;
  username?: string;
  password?: string;
  table?: string;
  sysId?: string;
};

const STORAGE_KEY = 'sn-sdk-doc-placeholders';
const EVENT = 'sn-sdk-doc-placeholders-changed';

export const DEFAULT_DISPLAY = {
  instance: '{instance}',
  username: '{username}',
  password: '{password}',
  table: '{TableName}',
  sysId: '{SysID}',
} as const;

export function readStoredValues(): PlaceholderValues {
  if (typeof window === 'undefined') return {};
  try {
    return JSON.parse(window.localStorage.getItem(STORAGE_KEY) ?? '{}');
  } catch {
    return {};
  }
}

export function storeValues(values: PlaceholderValues): void {
  window.localStorage.setItem(STORAGE_KEY, JSON.stringify(values));
  window.dispatchEvent(new Event(EVENT));
}

// Re-renders the caller whenever stored values change (any tab component,
// this or another page section). SSR-safe: first render uses defaults.
export function usePlaceholderValues(): PlaceholderValues {
  const [values, setValues] = useState<PlaceholderValues>({});
  useEffect(() => {
    const sync = () => setValues(readStoredValues());
    sync();
    window.addEventListener(EVENT, sync);
    window.addEventListener('storage', sync);
    return () => {
      window.removeEventListener(EVENT, sync);
      window.removeEventListener('storage', sync);
    };
  }, []);
  return values;
}

export function substitutePlaceholders(code: string, values: PlaceholderValues): string {
  const v = {...DEFAULT_DISPLAY, ...pruneEmpty(values)};
  const tokens: Record<string, string> = {
    xSDK_USERNAMEx: v.username,
    xSDK_PASSWORDx: v.password,
    xSDK_SN_INSTANCEx: v.instance,
    xSDK_SN_URLx: `https://${v.instance}.service-now.com`,
    xSDK_SN_TABLEx: v.table,
    xSDK_SN_TABLE_SYS_IDx: v.sysId,
  };
  return code.replace(/xSDK_[A-Z_]+x/g, (t) => tokens[t] ?? t);
}

function pruneEmpty(values: PlaceholderValues): PlaceholderValues {
  return Object.fromEntries(
    Object.entries(values).filter(([, v]) => typeof v === 'string' && v.trim() !== ''),
  );
}
