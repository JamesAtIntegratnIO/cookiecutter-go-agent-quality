import re
import sys

project_slug = "{{ cookiecutter.project_slug }}"
module_path = "{{ cookiecutter.module_path }}"
package_name = "{{ cookiecutter.package_name }}"
coverage_threshold = "{{ cookiecutter.coverage_threshold }}"

slug_regex = re.compile(r"^[a-z][a-z0-9-]*$")
package_regex = re.compile(r"^[a-z][a-z0-9_]*$")
module_regex = re.compile(r"^[A-Za-z0-9_.\-/]+/[A-Za-z0-9_.\-/]+$")

errors = []

if not slug_regex.match(project_slug):
    errors.append("project_slug must use lowercase letters, numbers, and hyphens, and start with a letter")

if not package_regex.match(package_name):
    errors.append("package_name must be a valid lowercase Go package identifier")

if not module_regex.match(module_path):
    errors.append("module_path must look like a Go module path, for example github.com/org/repo")

try:
    threshold = int(coverage_threshold)
    if threshold < 0 or threshold > 100:
        errors.append("coverage_threshold must be between 0 and 100")
except ValueError:
    errors.append("coverage_threshold must be an integer")

if errors:
    for error in errors:
        print(f"ERROR: {error}", file=sys.stderr)
    sys.exit(1)
