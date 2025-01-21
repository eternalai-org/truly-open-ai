import os
import json
import pprint

OUTPUT_FOLDER = "artifacts/artifacts"
MODEL_ARTIFACT_DIR = "artifacts/contracts/FunctionalModel.sol"
MODEL_ARTIFACT_OUTPUT = "models"
LAYER_ARTIFACT_DIR = "artifacts/contracts/libraries/layers-new"
LAYER_ARTIFACT_OUTPUT = "layers"

EXCLUDE_FILES = {"ILayer", "Layers"}

def create_output_directory(output_path):
    os.makedirs(output_path, exist_ok=True)
    init_file = os.path.join(output_path, "__init__.py")
    with open(init_file, "w") as f:
        f.write("")

def export_artifact_to_python(input_path, output_path):
    with open(input_path, "r") as f:
        artifact_json = f.read()
    artifact_dict = pprint.pformat(json.loads(artifact_json), compact=False)
    variable_name = "CONTRACT_ARTIFACT"
    with open(output_path, "w") as f:
        f.write(f"{variable_name} = {artifact_dict}")

def process_artifacts(artifact_dir, output_subdir, exclude_files=None):
    create_output_directory(output_subdir)
    for root, _, files in os.walk(artifact_dir):
        for filename in files:
            if filename.endswith(".json") and not filename.endswith(".dbg.json"):
                name = filename.split('.')[0]
                if exclude_files and name in exclude_files:
                    continue
                artifact_path = os.path.join(root, filename)
                output_path = os.path.join(output_subdir, f"{name}.py")
                export_artifact_to_python(artifact_path, output_path)

if __name__ == "__main__":
    if not os.path.exists(OUTPUT_FOLDER):
        os.makedirs(OUTPUT_FOLDER)
    with open(os.path.join(OUTPUT_FOLDER, "__init__.py"), "w") as f:
        f.write("")
    process_artifacts(MODEL_ARTIFACT_DIR, os.path.join(OUTPUT_FOLDER, MODEL_ARTIFACT_OUTPUT))
    process_artifacts(LAYER_ARTIFACT_DIR, os.path.join(OUTPUT_FOLDER, LAYER_ARTIFACT_OUTPUT), EXCLUDE_FILES)