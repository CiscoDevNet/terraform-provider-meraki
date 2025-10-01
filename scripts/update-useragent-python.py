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
    """Actualiza el UserAgent en provider.go"""
    provider_file = 'internal/provider/provider.go'
    
    try:
        # Leer el archivo
        with open(provider_file, 'r') as f:
            lines = f.readlines()
        
        # Crear backup
        with open(provider_file + '.backup', 'w') as f:
            f.writelines(lines)
        
        # Buscar y reemplazar la línea UserAgent
        updated = False
        for i, line in enumerate(lines):
            if 'UserAgent' in line and '=' in line:
                # Reemplazar toda la línea con el nuevo UserAgent
                lines[i] = f'\tc.UserAgent = "MerakiTerraform/{version} Cisco"\n'
                updated = True
                # Eliminar líneas adicionales que puedan contener restos del UserAgent anterior
                j = i + 1
                while j < len(lines) and ('Cisco' in lines[j] or lines[j].strip() == ''):
                    if 'Cisco' in lines[j]:
                        lines.pop(j)
                    else:
                        j += 1
                break
        
        if not updated:
            # Si no se encontró UserAgent, insertar antes de "data := MerakiProviderData"
            for i, line in enumerate(lines):
                if 'data := MerakiProviderData' in line:
                    lines.insert(i, f'\tc.UserAgent = "MerakiTerraform/{version} Cisco"\n')
                    updated = True
                    break
        
        if updated:
            # Escribir el archivo actualizado
            with open(provider_file, 'w') as f:
                f.writelines(lines)
            
            print(f"✅ UserAgent actualizado exitosamente a: MerakiTerraform/{version} Cisco")
            return True
        else:
            print("❌ Error: No se pudo encontrar dónde insertar UserAgent")
            return False
            
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
