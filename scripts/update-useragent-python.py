#!/usr/bin/env python3

import sys
import re
import os

def extract_version_from_changelog():
    """Extracts the latest version from changelog"""
    try:
        with open('CHANGELOG.md', 'r') as f:
            content = f.read()
        
        # Search for first line starting with ## X.Y.Z
        match = re.search(r'^## ([0-9]+\.[0-9]+\.[0-9]+)', content, re.MULTILINE)
        if match:
            version = match.group(1)
            return version
        else:
            print("Error: No version found in changelog")
            return None
    except FileNotFoundError:
        print("Error: CHANGELOG.md not found")
        return None

def update_useragent(version):
    """Updates UserAgent in provider.go using regex to avoid formatting issues"""
    provider_file = 'internal/provider/provider.go'
    
    try:
        # Read complete file
        with open(provider_file, 'r') as f:
            content = f.read()
        
        # Create backup
        with open(provider_file + '.backup', 'w') as f:
            f.write(content)
        
        # Regex pattern to find any line containing UserAgent
        # This handles cases where the string might be split across multiple lines
        useragent_pattern = r'c\.UserAgent\s*=\s*"[^"]*"(?:\s*"[^"]*")*'
        
        # Check if UserAgent already exists
        if re.search(useragent_pattern, content, re.MULTILINE):
            # Replace any line containing UserAgent with new format
            new_content = re.sub(
                useragent_pattern,
                f'c.UserAgent = "MerakiTerraform/{version} Cisco"',
                content,
                flags=re.MULTILINE
            )
        else:
            # If UserAgent doesn't exist, insert it before "data := MerakiProviderData"
            if 'data := MerakiProviderData' in content:
                new_content = content.replace(
                    'data := MerakiProviderData',
                    f'\tc.UserAgent = "MerakiTerraform/{version} Cisco"\n\n\tdata := MerakiProviderData'
                )
            else:
                print("❌ Error: Could not find where to insert UserAgent")
                return False
        
        # Write updated file
        with open(provider_file, 'w') as f:
            f.write(new_content)
        
        print(f"✅ UserAgent successfully updated to: MerakiTerraform/{version} Cisco")
        return True
            
    except Exception as e:
        print(f"❌ Error: {e}")
        return False

def main():
    # Get version from argument or changelog
    if len(sys.argv) > 1:
        version = sys.argv[1]
        print(f"Using provided version: {version}")
    else:
        print("Extracting version from changelog...")
        version = extract_version_from_changelog()
        if not version:
            sys.exit(1)
        print(f"Version found: {version}")
    
    # Update UserAgent
    if update_useragent(version):
        print("✅ Update completed")
    else:
        print("❌ Update error")
        sys.exit(1)

if __name__ == "__main__":
    main()