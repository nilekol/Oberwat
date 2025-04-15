#!/bin/bash

# Update package list
sudo apt-get update

# Check if Nginx is installed
if ! which nginx > /dev/null; then
    echo "Nginx not found, installing..."
    sudo apt-get install -y nginx
else
    echo "Nginx is already installed."
fi

# Copy Nginx config from the repo to the proper directory
sudo cp nginx/oberwat.conf /etc/nginx/conf.d/

# Check Nginx config for errors
sudo nginx -t

# Reload Nginx to apply changes
sudo systemctl reload nginx

# Start Nginx if it's not running
sudo systemctl start nginx

# Enable Nginx to start on boot
sudo systemctl enable nginx

# Final confirmation message
echo "Nginx setup completed!"
