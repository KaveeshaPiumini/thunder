/**
 * Copyright (c) 2026, WSO2 LLC. (https://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/**
 * Mobile-only API reference — stack navigation pattern.
 *
 * Two full-screen views that fill the space below the existing ThunderID
 * navbar (no additional header needed):
 *
 *   List view  — search box + collapsible tag groups + endpoint rows.
 *                Selecting an endpoint pushes to the detail view.
 *   Detail view — back bar showing the tag name + full scrollable
 *                 endpoint detail (method, path, params, responses).
 *
 * Single scroll axis per screen; no competing split panels.
 */

import React, {useCallback, useEffect, useMemo, useState} from 'react';
import {parse} from 'yaml';

// ─── Types ────────────────────────────────────────────────────────────────────

interface ParamSchema {
  type?: string;
  format?: string;
  enum?: string[];
  $ref?: string;
}

interface Param {
  name: string;
  in: 'path' | 'query' | 'header' | 'cookie';
  required?: boolean;
  description?: string;
  schema?: ParamSchema;
}

interface ResponseObj {
  description?: string;
}

interface FlatOp {
  id: string;
  method: string;
  path: string;
  summary: string;
  description?: string;
  tags: string[];
  parameters: Param[];
  requestBody?: {
    required?: boolean;
    description?: string;
    content?: Record<string, unknown>;
  };
  responses: Record<string, ResponseObj>;
}

interface TagGroup {
  name: string;
  description: string;
  ops: FlatOp[];
}

// ─── Constants ────────────────────────────────────────────────────────────────

const HTTP_METHODS = ['get', 'post', 'put', 'delete', 'patch', 'head', 'options'];

const METHOD_COLOR: Record<string, string> = {
  delete: '#f93e3e',
  get: '#61affe',
  head: '#9012fe',
  options: '#0d5aa7',
  patch: '#50e3c2',
  post: '#49cc90',
  put: '#fca130',
};

const STATUS_COLOR: Record<string, string> = {
  '2': '#49cc90',
  '3': '#61affe',
  '4': '#fca130',
  '5': '#f93e3e',
};

const VALID_PARAM_LOCATIONS = new Set(['path', 'query', 'header', 'cookie']);

// ─── Spec helpers ─────────────────────────────────────────────────────────────

function buildTagGroups(spec: Record<string, unknown>): TagGroup[] {
  const specTagDefs = (spec.tags as Array<{name: string; description?: string}>) ?? [];
  const tagMap = new Map<string, TagGroup>();

  for (const t of specTagDefs) {
    tagMap.set(t.name, {description: t.description ?? '', name: t.name, ops: []});
  }

  const paths = (spec.paths as Record<string, Record<string, unknown>>) ?? {};

  for (const [path, pathItem] of Object.entries(paths)) {
    for (const method of HTTP_METHODS) {
      const rawOp = pathItem[method] as Record<string, unknown> | undefined;
      if (!rawOp) continue;

      const opTags = (rawOp.tags as string[] | undefined) ?? ['Other'];

      // Only include parameters that have a recognised 'in' location.
      // Parameters without a valid location would otherwise render as
      // an "UNDEFINED PARAMETERS" section.
      const parameters = ((rawOp.parameters as Param[]) ?? []).filter(
        p => p?.in && VALID_PARAM_LOCATIONS.has(p.in),
      );

      const op: FlatOp = {
        description: rawOp.description as string | undefined,
        id: `${method}:${path}`,
        method,
        parameters,
        path,
        requestBody: rawOp.requestBody as FlatOp['requestBody'],
        responses: (rawOp.responses as Record<string, ResponseObj>) ?? {},
        summary: (rawOp.summary as string) ?? path,
        tags: opTags,
      };

      for (const tag of opTags) {
        if (!tagMap.has(tag)) {
          tagMap.set(tag, {description: '', name: tag, ops: []});
        }
        tagMap.get(tag)!.ops.push(op);
      }
    }
  }

  return [...tagMap.values()].filter(g => g.ops.length > 0);
}

// ─── Shared atoms ─────────────────────────────────────────────────────────────

function MethodBadge({method}: {method: string}) {
  return (
    <span
      style={{
        background: METHOD_COLOR[method] ?? '#888',
        borderRadius: 3,
        color: '#fff',
        display: 'inline-block',
        flexShrink: 0,
        fontSize: '0.67rem',
        fontWeight: 700,
        letterSpacing: '0.03em',
        minWidth: 52,
        padding: '2px 5px',
        textAlign: 'center',
        textTransform: 'uppercase',
      }}
    >
      {method}
    </span>
  );
}

// ─── List view ────────────────────────────────────────────────────────────────

interface ListViewProps {
  filteredGroups: TagGroup[];
  expandedTags: Set<string>;
  search: string;
  selectedOpId: string | null;
  onSearch: (q: string) => void;
  onToggleTag: (name: string) => void;
  onSelectOp: (op: FlatOp) => void;
}

function ListView({
  filteredGroups,
  expandedTags,
  search,
  selectedOpId,
  onSearch,
  onToggleTag,
  onSelectOp,
}: ListViewProps) {
  return (
    <div style={{display: 'flex', flex: 1, flexDirection: 'column', minHeight: 0, overflow: 'hidden'}}>
      {/* Search — always visible at the top */}
      <div
        style={{
          background: 'var(--oxygen-palette-background-paper)',
          borderBottom: '1px solid var(--ifm-color-emphasis-200)',
          flexShrink: 0,
          padding: '10px 14px',
        }}
      >
        <input
          placeholder="Search endpoints…"
          style={{
            background: 'var(--ifm-background-color)',
            border: '1px solid var(--ifm-color-emphasis-300)',
            borderRadius: 6,
            boxSizing: 'border-box',
            color: 'var(--ifm-font-color-base)',
            fontSize: '0.9rem',
            outline: 'none',
            padding: '8px 10px',
            width: '100%',
          }}
          type="search"
          value={search}
          onChange={e => onSearch(e.target.value)}
        />
      </div>

      {/* Tag groups — full remaining height, scrollable */}
      <div
        style={{
          flex: 1,
          minHeight: 0,
          overflowY: 'auto',
          WebkitOverflowScrolling: 'touch',
          overscrollBehavior: 'contain',
        }}
      >
        {filteredGroups.map(group => {
          const isExpanded = expandedTags.has(group.name);
          return (
            <div key={group.name}>
              <button
                style={{
                  alignItems: 'center',
                  background: 'transparent',
                  border: 'none',
                  borderBottom: '1px solid var(--ifm-color-emphasis-200)',
                  color: 'var(--ifm-font-color-base)',
                  cursor: 'pointer',
                  display: 'flex',
                  gap: 10,
                  padding: '12px 14px',
                  textAlign: 'left',
                  width: '100%',
                }}
                onClick={() => onToggleTag(group.name)}
              >
                <span
                  style={{
                    display: 'inline-block',
                    flexShrink: 0,
                    fontSize: '0.6rem',
                    opacity: 0.4,
                    transform: isExpanded ? 'rotate(90deg)' : undefined,
                    transition: 'transform 0.15s ease',
                  }}
                >
                  ▶
                </span>
                <span style={{flex: 1, fontSize: '0.92rem', fontWeight: 600}}>
                  {group.name}
                </span>
                <span
                  style={{
                    fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                    fontSize: '0.72rem',
                    opacity: 0.35,
                  }}
                >
                  {group.ops.length}
                </span>
              </button>

              {isExpanded &&
                group.ops.map(op => {
                  const isActive = selectedOpId === op.id;
                  return (
                    <button
                      key={op.id}
                      style={{
                        alignItems: 'center',
                        background: isActive
                          ? 'color-mix(in srgb, var(--ifm-color-primary) 10%, transparent)'
                          : 'transparent',
                        border: 'none',
                        borderBottom: '1px solid var(--ifm-color-emphasis-100)',
                        borderLeft: isActive
                          ? '3px solid var(--ifm-color-primary)'
                          : '3px solid transparent',
                        color: 'var(--ifm-font-color-base)',
                        cursor: 'pointer',
                        display: 'flex',
                        gap: 10,
                        padding: '9px 14px 9px 22px',
                        textAlign: 'left',
                        width: '100%',
                      }}
                      onClick={() => onSelectOp(op)}
                    >
                      <MethodBadge method={op.method} />
                      <span
                        style={{
                          fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                          fontSize: '0.78rem',
                          opacity: isActive ? 1 : 0.7,
                          wordBreak: 'break-all',
                        }}
                      >
                        {op.path}
                      </span>
                    </button>
                  );
                })}
            </div>
          );
        })}

        {filteredGroups.length === 0 && search && (
          <p
            style={{
              fontSize: '0.88rem',
              margin: 0,
              opacity: 0.45,
              padding: '20px 14px',
              textAlign: 'center',
            }}
          >
            No endpoints match &ldquo;{search}&rdquo;
          </p>
        )}
      </div>
    </div>
  );
}

// ─── Detail view ──────────────────────────────────────────────────────────────

interface DetailViewProps {
  op: FlatOp;
  tagName: string;
  onBack: () => void;
}

function SectionLabel({children}: {children: React.ReactNode}) {
  return (
    <p
      style={{
        fontSize: '0.67rem',
        fontWeight: 700,
        letterSpacing: '0.07em',
        margin: '0 0 8px',
        opacity: 0.45,
        textTransform: 'uppercase',
      }}
    >
      {children}
    </p>
  );
}

function DetailView({op, tagName, onBack}: DetailViewProps) {
  const paramGroups = useMemo(() => {
    const groups: Record<string, Param[]> = {};
    for (const p of op.parameters) {
      (groups[p.in] ??= []).push(p);
    }
    return Object.entries(groups);
  }, [op.parameters]);

  const responses = useMemo(
    () => Object.entries(op.responses ?? {}),
    [op.responses],
  );

  const contentTypes = useMemo(
    () => (op.requestBody?.content ? Object.keys(op.requestBody.content) : []),
    [op.requestBody],
  );

  return (
    <div style={{display: 'flex', flex: 1, flexDirection: 'column', minHeight: 0, overflow: 'hidden'}}>
      {/* Back bar — the only chrome added here; ThunderID navbar is above. */}
      <button
        style={{
          alignItems: 'center',
          background: 'var(--oxygen-palette-background-paper)',
          border: 'none',
          borderBottom: '1px solid var(--ifm-color-emphasis-200)',
          color: 'var(--ifm-color-primary)',
          cursor: 'pointer',
          display: 'flex',
          flexShrink: 0,
          fontSize: '0.88rem',
          fontWeight: 600,
          gap: 6,
          padding: '12px 14px',
          textAlign: 'left',
          width: '100%',
        }}
        onClick={onBack}
      >
        <span style={{fontSize: '1rem', lineHeight: 1}}>←</span>
        {tagName}
      </button>

      {/* Endpoint content — full remaining height, scrollable */}
      <div
        style={{
          flex: 1,
          minHeight: 0,
          overflowY: 'auto',
          WebkitOverflowScrolling: 'touch',
          overscrollBehavior: 'contain',
          padding: '20px 16px 48px',
        }}
      >
        <div
          style={{
            alignItems: 'center',
            display: 'flex',
            flexWrap: 'wrap',
            gap: 8,
            marginBottom: 6,
          }}
        >
          <MethodBadge method={op.method} />
          <code style={{fontSize: '0.85rem', opacity: 0.85, wordBreak: 'break-all'}}>
            {op.path}
          </code>
        </div>

        <h2 style={{fontSize: '1.05rem', fontWeight: 700, margin: '0 0 10px'}}>
          {op.summary}
        </h2>

        {op.description && (
          <p style={{fontSize: '0.88rem', lineHeight: 1.6, marginBottom: 20, opacity: 0.72}}>
            {op.description}
          </p>
        )}

        {/* Parameters */}
        {paramGroups.map(([location, params]) => (
          <div key={location} style={{marginBottom: 20}}>
            <SectionLabel>{location} parameters</SectionLabel>
            <div
              style={{
                border: '1px solid var(--ifm-color-emphasis-200)',
                borderRadius: 6,
                overflow: 'hidden',
              }}
            >
              {params.map((p, i) => (
                <div
                  key={p.name}
                  style={{
                    borderTop: i > 0 ? '1px solid var(--ifm-color-emphasis-200)' : undefined,
                    padding: '9px 12px',
                  }}
                >
                  <div
                    style={{
                      alignItems: 'center',
                      display: 'flex',
                      flexWrap: 'wrap',
                      gap: '3px 8px',
                      marginBottom: p.description ? 3 : 0,
                    }}
                  >
                    <code style={{fontSize: '0.82rem', fontWeight: 600}}>{p.name}</code>
                    {p.required && (
                      <span
                        style={{
                          background: 'rgba(249,62,62,0.1)',
                          borderRadius: 3,
                          color: '#f93e3e',
                          fontSize: '0.67rem',
                          fontWeight: 700,
                          padding: '1px 5px',
                        }}
                      >
                        required
                      </span>
                    )}
                    {p.schema?.type && (
                      <span
                        style={{
                          fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                          fontSize: '0.72rem',
                          opacity: 0.5,
                        }}
                      >
                        {p.schema.type}
                        {p.schema.format ? ` (${p.schema.format})` : ''}
                      </span>
                    )}
                  </div>
                  {p.description && (
                    <p style={{fontSize: '0.82rem', lineHeight: 1.4, margin: 0, opacity: 0.62}}>
                      {p.description}
                    </p>
                  )}
                  {p.schema?.enum && (
                    <p
                      style={{
                        fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                        fontSize: '0.72rem',
                        margin: '3px 0 0',
                        opacity: 0.5,
                      }}
                    >
                      Enum: {p.schema.enum.join(', ')}
                    </p>
                  )}
                </div>
              ))}
            </div>
          </div>
        ))}

        {/* Request body */}
        {op.requestBody && (
          <div style={{marginBottom: 20}}>
            <SectionLabel>
              request body{op.requestBody.required ? ' (required)' : ' (optional)'}
            </SectionLabel>
            <div
              style={{
                border: '1px solid var(--ifm-color-emphasis-200)',
                borderRadius: 6,
                overflow: 'hidden',
                padding: '9px 12px',
              }}
            >
              {contentTypes.length > 0 && (
                <p
                  style={{
                    fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                    fontSize: '0.78rem',
                    margin: 0,
                    opacity: 0.55,
                  }}
                >
                  {contentTypes.join(', ')}
                </p>
              )}
              {op.requestBody.description && (
                <p
                  style={{
                    fontSize: '0.85rem',
                    margin: contentTypes.length ? '4px 0 0' : 0,
                    opacity: 0.7,
                  }}
                >
                  {op.requestBody.description}
                </p>
              )}
            </div>
          </div>
        )}

        {/* Responses */}
        {responses.length > 0 && (
          <div style={{marginBottom: 20}}>
            <SectionLabel>responses</SectionLabel>
            <div
              style={{
                border: '1px solid var(--ifm-color-emphasis-200)',
                borderRadius: 6,
                overflow: 'hidden',
              }}
            >
              {responses.map(([code, resp], i) => (
                <div
                  key={code}
                  style={{
                    alignItems: 'flex-start',
                    borderTop: i > 0 ? '1px solid var(--ifm-color-emphasis-200)' : undefined,
                    display: 'flex',
                    gap: 12,
                    padding: '10px 12px',
                  }}
                >
                  <span
                    style={{
                      color: STATUS_COLOR[code[0]] ?? 'inherit',
                      flexShrink: 0,
                      fontFamily: 'var(--ifm-font-family-monospace, monospace)',
                      fontSize: '0.82rem',
                      fontWeight: 700,
                      minWidth: 36,
                    }}
                  >
                    {code}
                  </span>
                  <span style={{fontSize: '0.85rem', lineHeight: 1.4, opacity: 0.72}}>
                    {resp.description ?? '—'}
                  </span>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

// ─── Main component ───────────────────────────────────────────────────────────

export interface MobileApiReferenceProps {
  specUrl: string;
}

export default function MobileApiReference({specUrl}: MobileApiReferenceProps) {
  const [tagGroups, setTagGroups] = useState<TagGroup[]>([]);
  const [loading, setLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  // List view state
  const [search, setSearch] = useState('');
  const [expandedTags, setExpandedTags] = useState<Set<string>>(new Set());

  // null = list view; set = detail view (stack navigation)
  const [selectedOp, setSelectedOp] = useState<FlatOp | null>(null);

  useEffect(() => {
    let cancelled = false;
    setLoading(true);
    setLoadError(null);

    fetch(specUrl)
      .then(r => {
        if (!r.ok) throw new Error(`HTTP ${r.status}`);
        return r.text();
      })
      .then(text => {
        if (cancelled) return;
        const spec = parse(text) as Record<string, unknown>;
        const groups = buildTagGroups(spec);
        setTagGroups(groups);
        // Land on the first endpoint immediately so the view is never empty.
        if (groups.length > 0 && groups[0].ops.length > 0) {
          setSelectedOp(groups[0].ops[0]);
          setExpandedTags(new Set([groups[0].name]));
        }
        setLoading(false);
      })
      .catch(err => {
        if (!cancelled) {
          setLoadError(String(err));
          setLoading(false);
        }
      });

    return () => {
      cancelled = true;
    };
  }, [specUrl]);

  const q = search.toLowerCase().trim();

  const filteredGroups = useMemo<TagGroup[]>(() => {
    if (!q) return tagGroups;
    return tagGroups
      .map(g => ({
        ...g,
        ops: g.ops.filter(
          op =>
            op.summary.toLowerCase().includes(q) ||
            op.path.toLowerCase().includes(q) ||
            g.name.toLowerCase().includes(q),
        ),
      }))
      .filter(g => g.ops.length > 0);
  }, [tagGroups, q]);

  const toggleTag = useCallback((name: string) => {
    setExpandedTags(prev => {
      const next = new Set(prev);
      next.has(name) ? next.delete(name) : next.add(name);
      return next;
    });
  }, []);

  const handleSelectOp = useCallback((op: FlatOp) => {
    setSelectedOp(op);
  }, []);

  const handleBack = useCallback(() => {
    setSelectedOp(null);
  }, []);

  const selectedTagName = useMemo(() => {
    if (!selectedOp) return '';
    return tagGroups.find(g => g.ops.some(o => o.id === selectedOp.id))?.name ?? '';
  }, [selectedOp, tagGroups]);

  // Rendered via a portal into document.body so that Docusaurus's mobile sidebar
  // transform (which creates a new stacking context) cannot break position: fixed.
  const rootStyle: React.CSSProperties = {
    background: 'var(--oxygen-palette-background-default)',
    bottom: 0,
    display: 'flex',
    flexDirection: 'column',
    left: 0,
    overflow: 'hidden',
    position: 'fixed',
    right: 0,
    top: 'calc(var(--ifm-navbar-height) + var(--docusaurus-announcement-bar-height, 0px))',
    zIndex: 100,
  };

  if (loading) {
    return (
      <div
        className="apis-page"
        style={{
          ...rootStyle,
          alignItems: 'center',
          display: 'flex',
          fontSize: '0.88rem',
          justifyContent: 'center',
          opacity: 0.45,
        }}
      >
        Loading API reference…
      </div>
    );
  }

  if (loadError) {
    return (
      <div
        className="apis-page"
        style={{...rootStyle, color: '#f93e3e', fontSize: '0.88rem', padding: 20}}
      >
        Failed to load API reference: {loadError}
      </div>
    );
  }

  return (
    <div className="apis-page" style={rootStyle}>
      {selectedOp ? (
        <DetailView onBack={handleBack} op={selectedOp} tagName={selectedTagName} />
      ) : (
        <ListView
          expandedTags={expandedTags}
          filteredGroups={filteredGroups}
          search={search}
          selectedOpId={null}
          onSearch={setSearch}
          onSelectOp={handleSelectOp}
          onToggleTag={toggleTag}
        />
      )}
    </div>
  );
}
