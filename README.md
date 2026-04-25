# ja2en

WezTerm運用前提の日本語→英語翻訳CLIツール。OpenRouter API 経由で叩いた瞬間に英訳が返る。

## 特徴

- **起動 ~10ms**: Go製シングルバイナリ
- **設定ファイル + プロファイル切替**: シンプル英訳とカスタムプロンプトを使い分け
- **クリップボード対応**: `--clip` で読み込み、`--paste` で書き戻し
- **APIキーは環境変数のみ**: 設定ファイルには書かない

## インストール

```bash
git clone <this-repo> ~/ja2en
cd ~/ja2en
make install        # $GOPATH/bin/ja2en に配置
```

## セットアップ

1. OpenRouter で API key を発行: <https://openrouter.ai/keys>
2. シェル設定に追記:
   ```bash
   export OPENROUTER_API_KEY="sk-or-..."
   ```
3. 設定ファイル生成:
   ```bash
   ja2en init
   ```

## 使い方

```bash
# 引数
ja2en "明日出社する"

# stdin
echo "今日は遅れる" | ja2en

# クリップボード読込
ja2en --clip

# 翻訳結果をクリップボードに書込
ja2en --paste "緊急対応します"

# 読込→翻訳→書戻し
ja2en --clip --paste

# プロファイル切替
ja2en --profile detailed "..."

# 一発でモデル指定
ja2en --model anthropic/claude-3-haiku "..."
```

## 設定

`~/.config/ja2en/config.toml`:

```toml
default_profile = "simple"
model = "openai/gpt-oss-120b:free"
api_base = "https://openrouter.ai/api/v1"
timeout_seconds = 30

[profiles.simple]
prompt = "..."
```

## ライセンス

MIT
