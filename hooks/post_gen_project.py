from pathlib import Path
import os

scripts = [
    Path("scripts/check-coverage.sh"),
    Path("scripts/quality.sh"),
    Path("scripts/agent-finalize.sh"),
    Path("scripts/check-acceptance.sh"),
]

for script in scripts:
    mode = script.stat().st_mode
    script.chmod(mode | 0o111)

print("Generated Go agent-quality repository.")
print("Next steps:")
print("  go mod tidy")
print("  make install-tools")
print("  make quality-fast")
print("  make agent-finalize")
