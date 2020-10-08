import fs from 'fs';
import toml from 'toml';

import { IConfiguration } from "./config";

function loadConfig(path: string): IConfiguration {
    const raw = fs.readFileSync(path).toString();
    return toml.parse(raw);
}

export default loadConfig;