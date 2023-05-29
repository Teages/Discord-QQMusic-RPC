from matplotlib import colors
from matplotlib.colors import hsv_to_rgb
import os
import subprocess

h_colors = 16
s_colors = 1
v_colors = 1

n_colors = h_colors * s_colors * v_colors

colors_list = [hsv_to_rgb((i / n_colors, 0.2863, 1)) for i in range(n_colors)]

hex_colors = [colors.rgb2hex(c) for c in colors_list]

with open('template.svg', 'r') as f:
    content = f.read()

if not os.path.exists('dist_svg'):
    os.mkdir('dist_svg')

for i, color in enumerate(hex_colors):
    new_content = content.replace('#606060', color)
    with open(f'dist_svg/cd_icon_{i}.svg', 'w') as f:
        f.write(new_content)

if not os.path.exists('dist_png'):
    os.mkdir('dist_png')

for i in range(n_colors):
    print(f"dist_svg/cd_icon_{i}.svg -> dist_png/cd_icon_{i}.png")
    subprocess.run(['inkscape', f'dist_svg/cd_icon_{i}.svg', f'--export-filename=dist_png/cd_icon_{i}.png'])
