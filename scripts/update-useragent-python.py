#!/usr/bin/env python3

import sys
import re
import os

def extract_version_from_changelog():
    """Extrae la última versión del changelog"""
    try:
        with open('CHANGELOG.md', 'r') as f:
            content = f.read()
        
        # Buscar la primera línea que empiece con ## X.Y.Z
        match = re.search(r'^## ([0-9]+\.[0-9]+\.[0-9]+)', content, re.MULTILINE)
        if match:
            version = match.group(1)
            return version
        else:
            print("Error: No se encontró versión en el changelog")
            return None
    except FileNotFoundError:
        print("Error: No se encontró CHANGELOG.md")
        return None

def update_useragent(version):
    """Actualiza el UserAgent en provider.go usando regex para evitar problemas de formato"""
    provider_file = 'internal/provider/provider.go'
    
    try:
        # Leer el archivo completo
        with open(provider_file, 'r') as f:
            content = f.read()
        
        # Crear backup
        with open(provider_file + '.backup', 'w') as f:
            f.write(content)
        
        # Patrón regex para encontrar cualquier línea que contenga UserAgent
        # Esto maneja casos donde el string puede estar dividido en múltiples líneas
        useragent_pattern = r'c\.UserAgent\s*=\s*"[^"]*"(?:\s*"[^"]*")*'
        
        # Buscar si ya existe UserAgent
        if re.search(useragent_pattern, content, re.MULTILINE):
            # Reemplazar cualquier línea que contenga UserAgent con el nuevo formato
            new_content = re.sub(
                useragent_pattern,
                f'c.UserAgent = "MerakiTerraform/{version} Cisco"',
                content,
                flags=re.MULTILINE
            )
        else:
            # Si no existe UserAgent, insertarlo antes de "data := MerakiProviderData"
            if 'data := MerakiProviderData' in content:
                new_content = content.replace(
                    'data := MerakiProviderData',
                    f'\tc.UserAgent = "MerakiTerraform/{version} Cisco"\n\n\tdata := MerakiProviderData'
                )
            else:
                print("❌ Error: No se pudo encontrar dónde insertar UserAgent")
                return False
        
        # Escribir el archivo actualizado
        with open(provider_file, 'w') as f:
            f.write(new_content)
        
        print(f"✅ UserAgent actualizado exitosamente a: MerakiTerraform/{version} Cisco")
        return True
            
    except Exception as e:
        print(f"❌ Error: {e}")
        return False

def main():
    # Obtener versión del argumento o del changelog
    if len(sys.argv) > 1:
        version = sys.argv[1]
        print(f"Usando versión proporcionada: {version}")
    else:
        print("Extrayendo versión del changelog...")
        version = extract_version_from_changelog()
        if not version:
            sys.exit(1)
        print(f"Versión encontrada: {version}")
    
    # Actualizar UserAgent
    if update_useragent(version):
        print("✅ Actualización completada")
    else:
        print("❌ Error en la actualización")
        sys.exit(1)

if __name__ == "__main__":
    main()