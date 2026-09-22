# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project state

This is a go-study-app/scratch TypeScript repository (see README.md: "go-study-app"). There is no application source code yet — only tooling and project scaffolding. When adding code, check with the user about intended structure before assuming conventions, since none exist yet.

## Commands

Package manager is pnpm (enforced via `.npmrc` `save-exact=true`; use exact versions when adding dependencies).

- `pnpm fmt` — format code with oxfmt
- `pnpm fmt:check` — check formatting without writing changes
- `pnpm tsc` — type-check with `tsc --noEmit` (no build step; this project is type-check only)

There are no test or lint scripts defined yet, and no build/run scripts — add them to `package.json` if the project grows beyond type-checking.

## Toolchain

- Node `24.21.0` and pnpm `12.5.1`, pinned via `.mise.toml` / `.node-version` — use `mise` if available.
- Formatting/linting is handled by **oxfmt** (`.oxfmtrc.json`: 120 print width, single quotes, sorted imports, sorted Tailwind classes, sorted `package.json`), not Prettier/ESLint.
- A `PostToolUse` hook in `.claude/settings.json` automatically runs `oxfmt` on any file edited or written via Claude Code — manual formatting of touched files is usually unnecessary.
- TypeScript config (`tsconfig.json`) targets ES2022/ESNext with bundler module resolution, `verbatimModuleSyntax`, and strict mode (`strict`, `noUnusedLocals`, `noUnusedParameters`, `noFallthroughCasesInSwitch`, `noUncheckedSideEffectImports`) — no emit, so this repo is not meant to compile to JS output on its own.

## Git conventions

- Commit messages are linted by commitlint (`commitlint.config.ts`, extending `@commitlint/config-conventional`) via a Husky `commit-msg` hook — use Conventional Commits format. `subject-case` and `body-max-line-length` rules are disabled.
- `git reset`, `git rebase`, `git rm`, and `sudo` are denied in Claude Code's permissions (`.claude/settings.json`) — do not attempt these.

## MCP servers

`playwright`, `context7`, and `serena` are configured in `.mcp.json` and enabled in `.claude/settings.json`.
